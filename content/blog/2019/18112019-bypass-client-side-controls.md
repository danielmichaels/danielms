+++
title = "Getting around client-side controls"
categories = ["web"]
tags = ["web", "security"]
slug = "hacking-client-side-controls"
date = "2019-12-13"
draft = "true"
+++

# Bypassing Client-Side Controls

There are several methods for bypassing client controls to gain further access into the application, or alter its intended actions.

## Hidden Form Fields

HTML forms can have tags that are hidden from the user i.e. not rendered by the browser. If you have done web development with Flask (and Flask-forms) you may have seen the hidden CSRF field.
These fields whilst not rendered by the browser can be altered by a non-regular user. There is two ways of doing this, downloading the source code, editing it and then reloading *your* source into the browser to effect the change.
Or, by using a proxy based tool (such as Burp) and intercepting the request and updated the response.

## HTTP Cookies

Another method is interception and alteration of cookies. Cookies are generally not seen by regular users and cannot be modified directly by those users. Using a proxy it would be possible to intercept and change the cookie response.

## URL Parameters

This is a common method used in websites to transmit data. For example, browsing a catalogue may look something like

```shell
http://shop.com/shop/?prod=3&pricecode=32
```

This could easily be altered by just changing URL in the browser bar. There is times when this is not possible such as:

- when embedded images are loaded using URLs containing parameters
- where URLs containing parameters are used to load a frame's contents
- where a form uses the `POST` method and its target URL contains preset parameters
- where an application uses pop-up windows or other techniques to conceal the browser location bar

Although, all of these are easily overcome using the right tool.

## Referer Header

Browsers will include the referrer header when submitting HTTP requests as it is used to indicate the URL from which the request originated. It is used as a mechanism for validating data that is transmitted from the client. Meaning that some developers will assume that clients will do the right thing and consider this a safe method for tracking or generating requests.

In the example of a user resetting their password, an application may check the `Referer` header to check that the requests are following the correct sequence of requests (ForgotPassword.ashx --> ResetPassword.ashx --> VerifyNewPassword.ashx etc etc).

### Hack Steps

1. Locate all instances in the application where hidden fields, cookies or URL parameters are being used to transmit data via the client.
2. Attempt to determine or guess the role that the item plays in the applications logic, based on the context in which it appears.
3. Modify the item's value in way that is relevant to its purpose in the app. Ascertain whether the app processes these arbitrary values, and whether this leaves it vulnerable to attack.

## Opaque Data

Sometimes data transmitted via the client will be encrypted or obfuscated in some manner. Hidden fields may still exist but the data it holds may be a string of random or garbage looking characters. In this case you can reasonably infer that there is some form of server-side validation in play. This could lead the application vulnerable as it is doing the validation on the backend.

## Capturing User Data: HTML Forms

Validating forms can be done on the front or backend. The following are examples of client-side validations that can become vulnerabilities:

### Length Limits

Hidden fields can also be used in error checking and validation of forms. Length limits are one such example. Intercepting and removing the validating hidden field will circumvent this protection.

### Hack Steps

1. Look for form data containing a `maxlength` attribute. Submit data that is longer than the maxlength.
2. If it is accepted, you can infer that the error is not replicated on the backend
3. This could be a sign of defective validation and lead to further attacks such as SQLi, XSS, or buffer overflows


## Script-Based Validation

Input validation that is done using scripts on the client-side (JS) are relatively simple and lack the granularity of backend validators. A form that accepts several fields which varying lengths and attributes (strings, passwords, int's, floats etc) may have some type of script based validation prior to submission.

```html
<!-- Example of a script-based validation -->
<form method="post" action="Shop.aspx?prod=2" onsumbit="return validateFrom(this)">
  ...snip...
</form>

<script>function valdateForm(theForm)
{
  var isInteger = /^\d+$/;
  var valid = isInteger.test(quantity) &&
      quantity > 0 && <= 50;
  if (!valid)
    alert('Please enter a valid quantity');
  return valid;
}
</script>
```

The `onsubmit` attribute instructs the browser to execute the `ValidateForm` function when the user clicks the submit button, and to submit only if the function returns True. This is good for the developer as it ensures that most users cannot submit incorrectly formatted data inside thier browser.
This can be circumvented most times. Turning off JavaScript in the browser is one method, which will make the `onsubmit` attribute redundant leading to form submission without the script functioning.
Although many applications rely heavily on JavaScript and turning it off my break the usability of the website entirely. Another option is to enter a benign (Known Good) value, intercept it via a proxy and then modify the values. A simple way of defeating JS validation checks.
Or, you can intercept the response and alter the scripts function entirely, such as making all submissions equal true in the above script.

### Hack Steps

1. Identify any cases where client side JS is used to validate input
2. Submit known good data to the server but alter it via a proxy
3. If it has HTML length limits, alter these and check if the server is doing any validation
4. **NOTE**: if their is multiple input fields that subject to client-side validation you must test each one individually by entering valid data in the fields you are not testing. If not, the server may stop processing the first invalid input it see's thus not processing any remaining input fields with invalid data which will give you false-negatives.

This method is very popular today as it reduces load on servers and increases the user experience by reducing load times and giving prompt feedback. But, this does not mean that all these sites are vulnerable. This attack relies on the sever not validating the input, and the attacker crafting data that causes some undesirable behaviour by the application. Basically, an application may be vulnerable to this attack but that does not mean it will provide a useful outcome in its exploitation.

## Disabled Elements

Sometimes there may be HTML elements that are greyed out, or disabled. Usually these elements have been set this way to prevent the user from altering the element, or possibly it has been left for testing purposes. When the form is processed the disabled element is not sent to the server. It is worth testing to see if this element is still able to be processed by the server - if so, seek to exploit this.

### Hack Steps

1. Look for disabled elements within each form of the application. For each one you find, try submitting it to the server along with the forms other parameters to determine the effect of doing so.
2. Disabled elements are not sent to the browser by default, so you will need to modify the HTML using a proxy service such as burp.


## Handling Client-Side Data Securely

All data received from the client is inherently untrustworthy.

### Transmitting Data via the Client

Many applications make the fundamental mistake of sending critical data such as product prices and discounts via the client in an unsafe manner.
In such an example, the server should hold the prices and only display them to the client - the client should not be required to submit the prices back to the server! Even if user accounts have different prices or discounts applied this can still be handled on the server side. This information could be handled in a per-user database, user profiles or session objects.
If developers decide (or knowingly accept this risk) they should encrypt or sign the data from the client.

Two common pitfalls should be avoided:

- Some methods of signing or encrypting data can be exploited via replay attacks. Eg, if the product price is encrypted before being stored in a hidden field, it may be possible to copy the encrypted price of a cheaper product and submit it in place of the original price. A mitigation could be to concatenate the product and price as a single item so that the string can be validated when an order is submitted that matches that product and price.
- If users know the plain text value of the encrypted strings they are sent, they may be able to carry out cryptographic attacks to discover the encrypted key the server is using. If successful they may then be able to encrypt arbitrary data back to server, thereby circumventing this security posture.

### Validating Client-Generated Data

By virtue of it being on the client, data generated their cannot in principle be validated securely.

- Lightweight client-side controls such as HTML form fields and JS can be circumvented easily and provide no assurance about the input the server receives
-
