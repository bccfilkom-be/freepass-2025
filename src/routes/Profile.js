const express = require("express");
const { ProfileController } = require('./../controllers');

const Profile = express.Router();

Profile.patch('/', ProfileController.Update);

module.exports = Profile