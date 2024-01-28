# DigitalOcean Space Uploader Example by Golang

This project serves as a practical example to demonstrate how to upload objects to DigitalOcean Spaces using Golang. It provides a simple yet effective way to interact with DigitalOcean's object storage service, allowing users to securely store and manage their digital assets. The project illustrates essential features such as setting up credentials, configuring access permissions, and performing upload operations using Go's AWS SDK package. Ideal for developers looking to integrate DigitalOcean Spaces into their Golang applications.

## Configuration for DigitalOcean Spaces

This project uses DigitalOcean Spaces for object storage. Follow the steps below to configure your DigitalOcean Spaces access.

### Prerequisites

Ensure you have the following before you start:

- DigitalOcean account
- Access and secret keys for DigitalOcean Spaces
- A bucket created in DigitalOcean Spaces

### Setup

1. **Set your DigitalOcean Spaces credentials and bucket details:**

   Replace the placeholders in your configuration file or environment variables with your actual DigitalOcean Spaces details:

   - `<DO_ACCESS_KEY>`: Your DigitalOcean Spaces access key.
   - `<DO_SECRET_KEY>`: Your DigitalOcean Spaces secret key.
   - `<REGION>`: The region where your DigitalOcean Space is hosted (e.g., `nyc3`).
   - `<BUCKET_NAME>`: The name of your DigitalOcean Space bucket.

2. **Access Control List (ACL) settings:**

   This project supports two types of ACL settings for uploaded objects:

   - `public-read`: Use this setting if you want the uploaded objects to be publicly accessible.
   - `private`: Use this if you want to keep the uploaded objects private.

   Set the ACL according to your preference in the provided configuration section of the code.

