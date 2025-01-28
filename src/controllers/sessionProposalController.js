const { SessionProposal } = require("../models");
const { Op } = require("sequelize");

exports.createProposal = async (req, res) => {
  const {
    body: { userId, title, description, startTime, endTime },
  } = req;

  try {
    const existingProposal = await SessionProposal.findOne({
      where: {
        user_id: userId,
        [Op.or]: [
          { start_time: { [Op.between]: [startTime, endTime] } },
          { end_time: { [Op.between]: [startTime, endTime] } },
        ],
      },
    });

    if (existingProposal) {
      return res
        .status(400)
        .json({ message: "You already have a proposal in this time period" });
    }

    const newProposal = await SessionProposal.create({
      user_id: userId,
      title,
      description,
      start_time: startTime,
      end_time: endTime,
    });
    res
      .status(201)
      .json({ message: "Proposal created successfully", newProposal });
  } catch (err) {
    res.status(500).json({ message: "Failed to create proposal" });
  }
};

exports.editProposal = async (req, res) => {
  const {
    params: { id },
    body: { userId, title, description, startTime, endTime },
  } = req;

  try {
    const proposal = await SessionProposal.findByPk(id);

    if (!proposal) {
      return res.status(404).json({ message: "Proposal not found" });
    }

    if (proposal.user_id !== userId) {
      return res
        .status(403)
        .json({ message: "You are not authorized to edit this proposal" });
    }

    await proposal.update({
      title,
      description,
      start_time: startTime,
      end_time: endTime,
    });
    res
      .status(200)
      .json({ message: "Proposal updated successfully", proposal });
  } catch (err) {
    res.status(500).json({ message: "Failed to edit proposal" });
  }
};

exports.deleteProposal = async (req, res) => {
  const {
    params: { id },
    body: { userId },
  } = req;

  try {
    const proposal = await SessionProposal.findByPk(id);

    if (!proposal) {
      return res.status(404).json({ message: "Proposal not found" });
    }

    if (proposal.user_id !== userId) {
      return res
        .status(403)
        .json({ message: "You are not authorized to delete this proposal" });
    }

    await proposal.destroy();
    res.status(200).json({ message: "Proposal deleted successfully" });
  } catch (err) {
    res.status(500).json({ message: "Failed to delete proposal" });
  }
};

exports.getAllProposals = async (_req, res) => {
  try {
    const proposals = await SessionProposal.findAll({
      attributes: { exclude: ["createdAt", "updatedAt"] },
    });
    res.status(200).json(proposals);
  } catch (err) {
    res.status(500).json({ message: "Failed to retrieve proposals" });
  }
};
