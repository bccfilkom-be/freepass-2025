const { Router } = require("express");
const {
  editProfile,
  viewProfile,
} = require("../controllers/userController.js");
const { authMiddleware } = require("../middlewares/authMiddleware.js");
const {
  createUserValidationSchema,
} = require("../validations/createUserValidationSchema.js");
const { checkSchema } = require("express-validator");

const router = Router();

router.patch(
  "/edit-profile/:id",
  authMiddleware,
  checkSchema(createUserValidationSchema(true)),
  editProfile
);

router.get("/profile/:id", authMiddleware, viewProfile);

module.exports = router;
