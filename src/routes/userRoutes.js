const { Router } = require("express");
const { editProfile } = require("../controllers/userController.js");
const authMiddleware = require("../middlewares/authMiddleware.js");
const {
  updateUserValidationSchema,
} = require("../validations/userValidationSchema.js");
const { checkSchema } = require("express-validator");

const router = Router();

router.patch(
  "/edit-profile/:id",
  authMiddleware,
  checkSchema(updateUserValidationSchema),
  editProfile
);

module.exports = router;
