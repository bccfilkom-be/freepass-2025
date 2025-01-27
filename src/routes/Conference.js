const express = require("express");
const { ConferenceController } = require('../controllers');

const Conference = express.Router();

Conference.get('/', ConferenceController.View)
Conference.post('/', ConferenceController.Store)
Conference.patch('/', ConferenceController.Update)
Conference.delete('/', ConferenceController.Delete)
module.exports = Conference