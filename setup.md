## Setup a Kubernetes Cluster in AWS environment using Kops, Cloud9.
Tools needed:
- AWS account
- Cloud 9 IDE setup
- kops, kubectl, jq

Setup Cloud9 IDE
- Follow this link for a video tutorial on how to setup Cloud 9 IDE in your AWS Account
- [https://youtu.be/jKOtDkeVjmY](https://youtu.be/jKOtDkeVjmY)

Create AWS_KEY , AWS_SECRET
- Follow this link for a video reference on how to create a demo user with Programmatic access keys.
- <<TBA>>

Configure aws-cli
```
aws configure
AWS Access Key ID [None]: <aws_key>
AWS Secret Access Key [None]: <aws_secret>
Default region name [None]: us-east-2
Default output format [None]:
```

Install kubectl
```
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl; sudo mv ./kubectl /usr/local/bin/kubectl

```

Install KOPS
```
curl -LO https://github.com/kubernetes/kops/releases/download/$(curl -s https://api.github.com/repos/kubernetes/kops/releases/latest | grep tag_name | cut -d '"' -f 4)/kops-linux-amd64
chmod +x kops-linux-amd64; sudo mv kops-linux-amd64 /usr/local/bin/kops
```

Verify KOPS
```
kops version
```

Install jq - a tool to parse JSON output in command line.
```
sudo yum install jq
```

Create a S3 Bucket
** Info: bucket name = news-kops-state-store **


```
export bucketname=news-kops-state-store
```

```
aws s3api create-bucket --bucket $bucketname --create-bucket-configuration LocationConstraint=us-east-2
```

```
bala:~/environment $ aws s3api create-bucket --bucket news-kops-state-store --create-bucket-configuration LocationConstraint=us-east-2                  
{
    "Location": "http://news-kops-state-store.s3.amazonaws.com/"
}
```

```
aws s3api put-bucket-versioning --bucket $bucketname  --versioning-configuration Status=Enabled
```

Setup Cluster Name

```
export KOPS_CLUSTER_NAME=news.k8s.local
export KOPS_STATE_STORE=s3://news-kops-state-store
```

Generate a key-pair
```
bala:~/environment $ ssh-keygen 
Generating public/private rsa key pair.
Enter file in which to save the key (/home/ec2-user/.ssh/id_rsa): 
Enter passphrase (empty for no passphrase): 
Enter same passphrase again: 
Your identification has been saved in /home/ec2-user/.ssh/id_rsa.
Your public key has been saved in /home/ec2-user/.ssh/id_rsa.pub.
The key fingerprint is:
SHA256:ym+ScaTtmBCKFrERcHOUVcILaRcnKMpQOD+sZyiqFuM ec2-user@ip-172-31-13-63
The key's randomart image is:
+---[RSA 2048]----+
|o+=.o*=oo        |
|++.o* o+         |
|o=+o o .         |
|.++ . . .        |
| +.o . +S        |
|=++ ..o.o        |
|=oo  .oB         |
|.E    =.o        |
|+      o.        |
+----[SHA256]-----+
bala:~/environment $ ls /home/ec2-user/.ssh/
authorized_keys  id_rsa  id_rsa.pub

```

Create a secret for k8s cluster
```
kops create secret --name news.k8s.local sshpublickey admin -i ~/.ssh/id_rsa.pub
```


Setup a cluster
```
kops create cluster --name=news.k8s.local --state=s3://news-kops-state-store --node-count=2 --node-size=t2.micro --master-size=t2.micro --zones=us-east-2a
```

This step will take some time. Probably 5-10 mins. You can use the follow two commands to check / monitor the progress of the cluster creation.

Verify the cluster creation
```
kops validate cluster
```

You can actually watch and see the progress in realtime.
```
watch kops validate cluster
```
To exit , press <CTRL+D> or <CTRL+C>


Verify kubelet config
```
kubectl version
kubectl cluster-info
```

```
kubectl cluster-info
Kubernetes master is running at https://api-news-k8s-local-2506fm-662440992.us-east-2.elb.amazonaws.com
KubeDNS is running at https://api-news-k8s-local-2506fm-662440992.us-east-2.elb.amazonaws.com/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

To access the cluster dashboard, construct the url as like follows.
{master url}/ui

Example:
https://api-news-k8s-local-2506fm-662440992.us-east-2.elb.amazonaws.com/ui


Lets try to login to the Cluster
Username: admin

Password:
To get the admin password

```
kops get secrets kube --type secret -oplaintext
```

Try to login. If it ask for token, use the following command to get the token and access the cluster.
```
kops get secrets admin --type secret -oplaintext
```

<!-- kubectl expose deployment nginx --type=LoadBalancer --name=nginx --port 80 -->

Mean while you can also verify your cluster with `kubectl`
Get Nodes
```
kubectl get nodes
```

Get Pods
```
kubectl get pods
```

Get Namespaces
```
kubectl get ns
```

## Optionally, you can try to do a simple nginx deployment.

```
kubectl run nginx --image=nginx --replicas=1
```

```
kubectl get pods
```

```
kubectl get deployment
```


```
kubectl expose deployment nginx --type=LoadBalancer --name=nginx --port 80
```
Note this step creates a LoadBalancer in your aws account. Please make sure you delete this service when you delete your cluster. 

```
bala:~/environment $ kubectl get svc
NAME         TYPE           CLUSTER-IP       EXTERNAL-IP                                                               PORT(S)        AGE
kubernetes   ClusterIP      100.64.0.1       <none>                                                                    443/TCP        39m
nginx        LoadBalancer   100.66.125.181   a74730dc5a92d11e8bfc80285e348f77-1201292986.us-east-2.elb.amazonaws.com   80:32552/TCP   3m
```

- Finally, Access the application on the External-IP **
    - curl a74730dc5a92d11e8bfc80285e348f77-1201292986.us-east-2.elb.amazonaws.com **

- Open your browser and test your application
    - a74730dc5a92d11e8bfc80285e348f77-1201292986.us-east-2.elb.amazonaws.com


Lets destroy the Cluster to save money.

Before destroying the cluster, lets delete those LoadBalancers that we created while exposing 
some k8s service.
```
kubectl get svc
```
Lets delete the service that we created previously: `nginx`

```
kubectl delete svc nginx
```


Now delete the cluster using KOPS.
Kops will clean all the resources created for our k8s cluster. This will take about 3-5 mins.
```
kops delete cluster --name ${KOPS_CLUSTER_NAME} --yes
```
This will take some time (2 mins)
```
kops delete cluster --name ${KOPS_CLUSTER_NAME} --yes
TYPE                    NAME                                                                            ID
autoscaling-config      master-us-east-2a.masters.news.k8s.local-20180826120309                         master-us-east-2a.masters.news.k8s.local-20180826120309
autoscaling-config      nodes.news.k8s.local-20180826120309                                             nodes.news.k8s.local-20180826120309
autoscaling-group       master-us-east-2a.masters.news.k8s.local                                        master-us-east-2a.masters.news.k8s.local
autoscaling-group       nodes.news.k8s.local                                                            nodes.news.k8s.local
dhcp-options            news.k8s.local                                                                  dopt-03f00a43c9c9082dd
iam-instance-profile    masters.news.k8s.local                                                          masters.news.k8s.local
iam-instance-profile    nodes.news.k8s.local                                                            nodes.news.k8s.local
iam-role                masters.news.k8s.local                                                          masters.news.k8s.local
iam-role                nodes.news.k8s.local                                                            nodes.news.k8s.local
instance                master-us-east-2a.masters.news.k8s.local                                        i-0dc880456c2563bcb
instance                nodes.news.k8s.local                                                            i-054df6b4989230a7f
instance                nodes.news.k8s.local                                                            i-090ec89878374bfc3
internet-gateway        news.k8s.local                                                                  igw-0f7576cc16a7487ea
keypair                 kubernetes.news.k8s.local-36:e7:90:81:62:f8:a7:d1:ae:8f:4d:bf:77:55:b5:5c       kubernetes.news.k8s.local-36:e7:90:81:62:f8:a7:d1:ae:8f:4d:bf:77:55:b5:5c
load-balancer           api.news.k8s.local                                                              api-news-k8s-local-2506fm
route-table             news.k8s.local                                                                  rtb-0aae3742d8b3368b1
security-group          api-elb.news.k8s.local                                                          sg-0e8519cb4a6f2a79e
security-group          masters.news.k8s.local                                                          sg-050dcaad4291cf23e
security-group          nodes.news.k8s.local                                                            sg-08d3d9441cd3da85f
subnet                  us-east-2a.news.k8s.local                                                       subnet-045e4fe604a3a17a5
volume                  a.etcd-events.news.k8s.local                                                    vol-003175f7cbe7edcd1
volume                  a.etcd-main.news.k8s.local                                                      vol-0ae5a354c23d9f373
vpc                     news.k8s.local                                                                  vpc-0c092f4c71aad8099


....
....

Not all resources deleted; waiting before reattempting deletion
        security-group:sg-08d3d9441cd3da85f
        dhcp-options:dopt-03f00a43c9c9082dd
        route-table:rtb-0aae3742d8b3368b1
        subnet:subnet-045e4fe604a3a17a5
        vpc:vpc-0c092f4c71aad8099
subnet:subnet-045e4fe604a3a17a5 ok
security-group:sg-08d3d9441cd3da85f     ok
route-table:rtb-0aae3742d8b3368b1       ok
vpc:vpc-0c092f4c71aad8099       ok
dhcp-options:dopt-03f00a43c9c9082dd     ok
Deleted kubectl config for news.k8s.local

Deleted cluster: "news.k8s.local"
```

- Lets delete the S3 Bucket also.

Make sure your bucketname is same. You can try this `echo $bucketname`.

Now lets delete the S3 bucket using the following snippet.
```
curl -O https://gist.githubusercontent.com/weavenet/f40b09847ac17dd99d16/raw/e9fad5e2cd16f6f54446acdd79d47212f178ac6b/delete_all_object_versions.sh

sh delete_all_object_versions.sh $bucketname;
aws s3 rb $KOPS_STATE_STORE
```

Thats all you have successfully installed a cluster