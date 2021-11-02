# Sezzle Online Shopping portal

## App Domain URI: <http://ec2-65-2-71-29.ap-south-1.compute.amazonaws.com:3000/register>

## GitHub Repository: <https://github.com/DeepikaAzad/shopping-portal>

System Requirements: Mysql 5.7, go1.16.9 darwin/amd64, node v14.18.1, yarn 1.22.17

To access the API via Postman:

- Register user
- Login User, Once user is logged-in, copy the `token` from the response and paste it into furthure API's header.


[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/b48b8cb0a82f30e99df3?action=collection%2Fimport)

Features:

1. Have implemented secure jwt login.
2. User can register.
3. For each logged-in user, user will be able to add item to cart, place the order, see the cart list, item list and see the order history .
4. Only authorized user can create item and see the user list.
5. Item Creation is not user specific. Other user can see the list of items which is created by other user (admin). Since role based access is not handled so user can create the items.
6. Cart will create on adding first itme select and cart will remove if user remove all item from cart.
7. All the APIs are securly oauth protected and data sharing between the Front-End and Back-End is in JSON.
8. Code is very modular, following the standard pratices with basic documentaion.

Technology:

- Golang
- Mysql
- React

Known Issues:

- Very Basic UI implemented.
- To see the User list and add item role based authorization not implemented.

Please feel free to contact me for any queries.

Deepika Azad
azaddeepika05@gmail.com
+91 9783839582
