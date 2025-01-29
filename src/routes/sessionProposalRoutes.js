const { Router } = require("express");
const {
  createProposal,
  editProposal,
  deleteProposal,
  getAllProposals,
  acceptProposal,
  rejectProposal,
} = require("../controllers/sessionProposalController.js");
const { authMiddleware } = require("../middlewares/authMiddleware.js");
const { authorizeRole } = require("../middlewares/roleMiddleware.js");
const { checkSchema } = require("express-validator");
const {
  sessionProposalSchema,
  idSchema,
} = require("../validations/createSessionProposalValidationSchema.js");

const router = Router();

router.get("/", authMiddleware, authorizeRole("coordinator"), getAllProposals);

router.post(
  "/",
  authMiddleware,
  checkSchema(sessionProposalSchema),
  createProposal
);

router.put(
  "/:id",
  authMiddleware,
  checkSchema({ ...sessionProposalSchema, ...idSchema }),
  editProposal
);

router.delete(
  "/:id",
  authMiddleware,
  checkSchema(idSchema(false)),
  deleteProposal
);

router.put(
  "/:id/accept",
  authMiddleware,
  authorizeRole("coordinator"),
  checkSchema(idSchema()),
  acceptProposal
);

router.put(
  "/:id/reject",
  authMiddleware,
  authorizeRole("coordinator"),
  checkSchema(idSchema()),
  rejectProposal
);

module.exports = router;
