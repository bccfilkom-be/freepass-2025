const { Router } = require("express");
const {
  createProposal,
  editProposal,
  deleteProposal,
} = require("../controllers/sessionProposalController.js");

const router = Router();

router.post("/", createProposal);
router.put("/:id", editProposal);
router.delete("/:id", deleteProposal);

module.exports = router;
