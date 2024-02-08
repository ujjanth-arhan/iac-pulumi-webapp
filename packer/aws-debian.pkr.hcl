packer {
  required_plugins {
    amazon = {
      source  = "github.com/hashicorp/amazon"
      version = ">= 1.0.0"
    }
  }
}

variable "source_ami" {
  type        = string
  description = "The AMI ID"
}

variable "ssh_username" {
  type        = string
  description = "Username to SSH into the instance"
}

variable "subnet_id" {
  type        = string
  description = "Subnet Id"
}

source "amazon-ebs" "debian" {
  ami_name        = "ama_${formatdate("YYYY_MM_DD_hh_mm_ss", timestamp())}"
  ami_description = "AMI for cloud class"

  instance_type = "t2.micro"
  source_ami    = "${var.source_ami}"
  ssh_username  = "${var.ssh_username}"
  subnet_id     = "${var.subnet_id}"


  ami_users = ["453138487766", "039149918584"]

  launch_block_device_mappings {
    delete_on_termination = true
    device_name           = "/dev/xvda"
    volume_size           = 8
    volume_type           = "gp2"
  }
}

build {
  name    = "debian-build"
  sources = ["source.amazon-ebs.debian"]
  provisioner "file" {
    source      = "./builds/"
    destination = "/tmp"
  }
  provisioner "shell" {
    environment_vars = [
      "DEBIAN_FRONTEND=noninteractive",
      "CHECKPOINT_DISABLE=1",
    ]
    script = "./packer/script.sh"
  }
}

