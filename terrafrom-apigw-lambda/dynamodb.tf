
# Create a DynamoDB instance
resource "aws_dynamodb_table" "table" {
  name           = var.db_table_name
  hash_key       = "id"
  read_capacity  = 5
  write_capacity = 5

  attribute {
    name = "id"
    type = "S"
  }
}




















