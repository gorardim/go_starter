package aws_ses

const VerifyEmailTemplate string = `
<!DOCTYPE html>
<html>
<head>
    <title>inTrips</title>
    <style type="text/css">
        @media only screen and (max-width: 600px) {
            h1 {
                font-size: 42px !important;
            }
            h3 {
                font-size: 16px !important;
            }
            p {
                font-size: 14px !important;
            }
        }
    </style>
</head>
<body style="font-family: 'Roboto', 'Helvetica', sans-serif; margin: 0; padding: 0;">

    <table align="center" cellpadding="0" cellspacing="0" style="margin: 0 auto;">
        <tr>
            <td>
                <table cellpadding="0" cellspacing="0" width="100%">
                    <tr>
                        <td bgcolor="#003050" style="background-color: #003050; color: #DDE1E1; padding: 10px; text-align: center;">
                            <h1 style="margin: 0; padding: 0; font-size: 30px; font-weight: 700;">inTrips</h1>
                            <p style="margin: 0; padding: 0; font-size: 10px; font-weight: 300;">INVEST AND TRAVEL</p>
                        </td>
                    </tr>
                    <tr>
                        <td bgcolor="#E4EFFF" style="background-color: #E4EFFF; padding: 20px;">
                            <h3 style="font-size: 18px; font-weight: 500; color: #000;">Dear Customer,</h3>
                            <p style="font-size: 13px; color: #2B2B2B;">Your one-time password for inTrips is:</p>
                            <h1 style="font-size: 30px; font-weight: 500; color: #FFA800; text-align: center; padding: 10px;">@@@@@@</h1>
                            <p style="font-size: 15px; color: #000; font-weight: 400;">Kindly use this OTP to complete your operation within the next 5 minutes. For your security, please do not share this OTP with anyone.</p>
                            <p style="font-size: 15px; color: #000; font-weight: 400;">If you have not initiated this action or have any concerns about your account security, please contact our support team immediately at <span style="color: #FFA800;">[Your Support Email or Phone Number]</span>.</p>
                            <p style="font-size: 15px; color: #000; margin-top: 20px; font-weight: 400;">Thank you for your cooperation.</p>
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <img src="https://firebasestorage.googleapis.com/v0/b/orlogo-12567.appspot.com/o/7d9f2fdd897b253ac6f73032312c0a98.jpeg?alt=media&token=56323f08-9063-451f-974a-ff79743691e3&_gl=1*8w7r3i*_ga*MTkwOTE2NDg5NS4xNjk6MTY5ODM5NTQyNS4y.ljEuMTY5ODM5NTc4My41MS4wLjA" alt="inTrips" style="width: 100%; height: auto; display: block;">
                        </td>
                    </tr>
                    <tr style="background: #e4efff;">
                        <td style="border-bottom: 1px solid #B8D5FF; padding-top: 20px; padding-bottom: 20px;"></td>
                    </tr>
                    <tr>
                        <td bgcolor="#E4EFFF" style="background-color: #E4EFFF; color: #5B5B5B; font-size: 10px !important; padding: 10px; text-align: center;">
                            <table width="100%" cellpadding="0" cellspacing="0" style="margin-top: 20px;">
                                <tr>
                                    <td style="text-align: right; padding-right: 20px;" width="50">© Copyright 2024.</td>
                                    <td width="50" style="text-align: left; border-left: 1px solid #959595; padding-left: 20px; vertical-align: top;">
                                        <div style="margin: 0; line-height: 1;">UNITED MASS TOURISM LLC.</div>
                                        <div style="margin: 0; line-height: 1;">All rights reserved.</div>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>
`
