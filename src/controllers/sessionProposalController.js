const { Session, SessionProposal } = require("../models");
const { Op } = require("sequelize");
const { validationResult, matchedData } = require("express-validator");

exports.createProposal = async (req, res) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() });
  }

  const { userId, title, description, startTime, endTime } = matchedData(req);

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
  } catch {
    res.status(500).json({ message: "Failed to create proposal" });
  }
};

exports.editProposal = async (req, res) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() });
  }

  const { id, userId, title, description, startTime, endTime } =
    matchedData(req);

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
  } catch {
    res.status(500).json({ message: "Failed to edit proposal" });
  }
};

exports.deleteProposal = async (req, res) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() });
  }

  const { id, userId } = matchedData(req);

  try {
    const proposal = await SessionProposal.findByPk(id);

    if (!proposal) {
      return res.status(404).json({ message: "Proposal not found" });
    }

    console.log(proposal.user_id, userId);

    if (proposal.user_id !== userId) {
      return res
        .status(403)
        .json({ message: "You are not authorized to delete this proposal" });
    }

    await proposal.destroy();
    res.status(200).json({ message: "Proposal deleted successfully" });
  } catch {
    res.status(500).json({ message: "Failed to delete proposal" });
  }
};

exports.getAllProposals = async (_req, res) => {
  try {
    const proposals = await SessionProposal.findAll({
      attributes: { exclude: ["createdAt", "updatedAt"] },
    });
    res.status(200).json(proposals);
  } catch {
    res.status(500).json({ message: "Failed to retrieve proposals" });
  }
};

exports.acceptProposal = async (req, res) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() });
  }

  const { id } = matchedData(req);

  try {
    const proposal = await SessionProposal.findByPk(id);

    if (!proposal) {
      return res.status(404).json({ message: "Proposal not found" });
    }

    await proposal.update({ status: "accepted" });

    const newSession = await Session.create({
      title: proposal.title,
      description: proposal.description,
      start_time: proposal.start_time,
      end_time: proposal.end_time,
      available_seats: proposal.available_seats || 20,
      user_id: proposal.user_id,
    });

    res
      .status(200)
      .json({ message: "Proposal accepted and session created", newSession });
  } catch {
    res
      .status(500)
      .json({ message: "Failed to accept proposal and create session" });
  }
};

exports.rejectProposal = async (req, res) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() });
  }

  const { id } = matchedData(req);

  try {
    const proposal = await SessionProposal.findByPk(id);

    if (!proposal) {
      return res.status(404).json({ message: "Proposal not found" });
    }

    await proposal.update({ status: "rejected" });

    res.status(200).json({ message: "Proposal rejected", proposal });
  } catch {
    res.status(500).json({ message: "Failed to reject proposal" });
  }
};
