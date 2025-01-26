const { Router } = require("express");
const { register, login } = require("../controllers/authController.js");
const {
  userValidationSchema,
} = require("../validations/createUserValidationSchema.js");
const { body, checkSchema } = require("express-validator");

const router = Router();

router.post("/register", checkSchema(userValidationSchema), register);
router.post(
  "/login",
  [
    body("email").isEmail().withMessage("Invalid email format"),
    body("password").notEmpty().withMessage("Password is required"),
  ],
  login
);

module.exports = router;
