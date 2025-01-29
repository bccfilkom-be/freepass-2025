const { Router } = require("express");
const {
  getAllSessions,
  leaveFeedback,
  deleteFeedback,
  registerForSession,
  editSession,
  deleteSession,
} = require("../controllers/sessionController.js");
const { authMiddleware } = require("../middlewares/authMiddleware.js");
const { checkSchema, param } = require("express-validator");
const {
  createSessionValidationSchema,
  leaveFeedbackSchema,
  registerForSessionSchema,
} = require("../validations/createSessionValidationSchema.js");

const router = Router();

router.get("/", getAllSessions);

router.post("/:id/feedback", checkSchema(leaveFeedbackSchema), leaveFeedback);

router.delete(
  "/:id/feedback/:feedbackId",
  [
    param("id").isInt().withMessage("Session ID must be an integer"),
    param("feedbackId").isInt().withMessage("Feedback ID must be an integer"),
  ],
  authMiddleware,
  deleteFeedback
);

router.post(
  "/:id/register",
  checkSchema(registerForSessionSchema),
  registerForSession
);

router.put(
  "/:id",
  checkSchema(createSessionValidationSchema(true)),
  editSession
);

router.delete(
  "/:id",
  [param("id").isInt().withMessage("Session ID must be an integer")],
  authMiddleware,
  deleteSession
);

module.exports = router;
