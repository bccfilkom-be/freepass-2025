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

const router = Router();

router.get("/", authMiddleware, authorizeRole("coordinator"), getAllProposals);
router.post("/", authMiddleware, createProposal);
router.put("/:id", authMiddleware, editProposal);
router.delete("/:id", authMiddleware, deleteProposal);
router.put(
  "/:id/accept",
  authMiddleware,
  authorizeRole("coordinator"),
  acceptProposal
);
router.put(
  "/:id/reject",
  authMiddleware,
  authorizeRole("coordinator"),
  rejectProposal
);

module.exports = router;
