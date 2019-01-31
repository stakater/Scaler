# Scaler

A tool to modify auto scaling groups.

## Usage

Scaler can be used to modify max, min and desired capacity of an auto scaling group.

Scaler supports the following flags

|Short Flag|Long Flag|Type|Description|
|----------|---------|----|-----------|
|-d | --desired    | int    | Desired no of instances (Required)|
|-h | --help       |        | Help for Scaler|
|-m | --max        | int    | Maximum no of instances (Required)|
|-i | --min        | int    | Minimum no of instances (Required)|
|-p | --provider   | string | Cloud provider to user. Valid Values (aws)|
|-r | --region     | string | Region in which auto scaling group exists (Required)|
|-a | --roleArn    | string | Arn of role to assume (Required)|
|-n | --scalerName | string | Name of Auto Scaling group (Required)|

## Run

Scale can be run by passing arguments to the published image. e.g

```docker run -it docker.io/stakater/scaler:0.0.1 --roleArn arn:aws:iam::449074299682:role/nodes.stackator.com --region us-west-2 --max 0 --min 0 --desired 0 --provider aws --scalerName nodes.stackator.com```


## Use Case

Scaler can be used in combination with Cronjob to control the cluster size during different times. e.g You can bring down your cloud cost by keeping the servers shut down on weekends

A sample configuration of kubernetes cronjob can be found [here](cronjob/example.yaml).

## Adding Support for Cloud Providers

[Provider](internal/pkg/providers/provider.go) can be implemented to modify auto scaling groups on other cloud providers e.g Azure, Google Cloud etc. If new parameters are needed then they can be added by modifying Scaler Options struct in [common.go](internal/pkg/cmd/common/common.go). The values will automatically be available in the Init method of new implementation of [Provider](internal/pkg/providers/provider.go)
