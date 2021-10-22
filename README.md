
https://docs.aws.amazon.com/ses/latest/DeveloperGuide/create-shared-credentials-file.html

my user had the SES required authorizations, but still I was getting "... is not authorized to perform `ses:SendEmail'" the cause was the MFA at account level, once disabled it worked as a charm :)