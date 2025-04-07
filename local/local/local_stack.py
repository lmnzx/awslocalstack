from aws_cdk import CfnOutput, Stack, aws_s3, aws_sqs
from constructs import Construct


class LocalStack(Stack):
    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)
        self.bucket = aws_s3.Bucket(self, "test-bucket")
        self.queue = aws_sqs.Queue(self, "my-custom-sqs-queue")

        CfnOutput(
            self,
            "BucketName",
            value=self.bucket.bucket_name,
            export_name="MyBucketName",
        )
        CfnOutput(
            self, "QueueUrl", value=self.queue.queue_url, export_name="MyQueueUrl"
        )
