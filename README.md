
# GCP OAuth Example
### TOC

- [Overviews](#-overview)
- [Enable Google APIs](#-enable-google-apis)
- [Setup The Project](#-setup-the-project)

###  Overview
This is a simple tutorial project demonstrating how to integrate Google OAuth using the GCP API.  
The project is structured with minimal layering for clarity and educational purposes. 

It demonstrates how to:

- Initialize Google OAuth with `client_id` and `client_secret`
- Handle OAuth redirects and callbacks
- Organize code with basic folder structure

> ⚠️ This project is not production-ready and is designed for learning and reference only.

### 🔧 Enable Google APIs

#### **Step 1**

To use this demo, please go to the [Google Cloud API Dashboard](https://console.cloud.google.com/apis/dashboard), select your project, and enable required APIs (e.g. OAuth 2.0, Drive API)

#### **Step 2**
Click "create project"
<img src="https://github.com/Soyuen/picture/blob/main/create_project.png?raw=true" alt="Create Project" width="500"/>

#### **Step 3**
Click **"Create Project"**.
Enter your desired project name and click **"Create"**.
<img src="https://github.com/Soyuen/picture/blob/main/new%20project.png?raw=true" alt="New Project" width="400"/>

#### **Step 4**
Click **“Credentials”** on the **left sidebar** under **APIs & Services**.
Enter your desired project name and click **"Create"**.
<img src="https://github.com/Soyuen/picture/blob/main/credentials.png?raw=true" alt="credentials" width="300"/>

#### **Step 5**
Then, click "Create Credentials" located near the top center of the page.
Next, click "OAuth Client ID" to create a new OAuth client.
<img src="https://github.com/Soyuen/picture/blob/main/create_credentials.png?raw=true" alt="create credentials" width="400"/>

#### **Step 6**
If you see a prompt to "Configure consent screen", please click on it to set up your OAuth consent details. And go back to the step 5.
If you don’t see it, you can skip this step.
<img src="https://github.com/Soyuen/picture/blob/main/consent_screen.png?raw=true" alt="consent_screen" width="600"/>

#### **Step 7**
Enter the required information, then click "Create". To create the credential.
You can get the client id and the client secret key.

### 🔧 Setup The Project

After you obtain the **Client ID** and **Client Secret** from Google Cloud Console, add the following variables to your `.env` file:
```env
GOOGLE_CLIENT_ID=XXXXXXXXX
GOOGLE_CLIENT_SECRET=XXXXXXXXXX
GOOGLE_REDIRECT_URL=XXXXXXXXXXXXX
```
```main
go run cmd/main.go
```
