const express = require("express");
const { AuthController } = require('./../controllers');

const Auth = express.Router();

Auth.post('/register', AuthController.Register);
Auth.post('/login', AuthController.Login);

module.exports = Auth