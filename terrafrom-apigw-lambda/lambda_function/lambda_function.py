import json

def lambda_handler(event, context):
    method = event['httpMethod']
    if method == 'GET':
        return {
            'statusCode': 200,
            'body': json.dumps('get items!')
        }
    if method == 'DELETE':
        # Delete item from DynamoDB
        return {
            'statusCode': 200,
            'body': json.dumps('Item deleted successfully')
        }
