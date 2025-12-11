# OAuth Merge Plan

## Overview

This document describes the plan for handling user registration with
email/password, OAuth login using Google/GitHub, and managing linked
providers in user settings.

## Database Design

### users

-   id
-   email
-   password_hash
-   firstname
-   lastname
-   avatar_url
-   email_verified
-   created_at
-   updated_at

### provider_accounts

-   id
-   user_id
-   provider
-   provider_user_id
-   access_token
-   refresh_token
-   raw_profile
-   created_at
-   updated_at

## Flow Cases

### 1. Register with Email/Password

-   Create base user.

### 2. Login with Email/Password

-   Standard login.

### 3. Existing Email User → Login with Google

-   Find user by email.
-   Create provider_accounts entry for Google.

### 4. Existing Email User → Login with GitHub

-   Similar to Google.

### 5. Provider with Different Email

-   Do NOT auto-merge.
-   Require user to link manually.

### 6. Link Provider from Admin Page

-   Redirect user to OAuth with `state=link`.
-   Backend links provider to logged-in user.

### 7. Unlink Provider

-   Cannot unlink last provider unless password exists.

### 8. First Time OAuth → Create User

-   Create user + provider_accounts.

### 9. Profile Sync Rules

-   Only sync firstname/lastname/avatar on first creation.
-   Never overwrite later.
