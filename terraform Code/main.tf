provider "aws" {
}
 
resource "aws_instance" "app_server" {
  ami           = "ami-04ff98ccbfa41c9ad"
  instance_type = "t2.micro"
 
  tags = {
    Name = "ExampleAppServerInstance"
  }
}

