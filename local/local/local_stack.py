from aws_cdk import (
    Stack,
    aws_s3,
    aws_sqs,
)
from constructs import Construct


class LocalStack(Stack):
    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)
        bucket = aws_s3.Bucket(self, "test-bucket")
        sqs = aws_sqs.Queue(self, "my-custom-sqs-queue")
