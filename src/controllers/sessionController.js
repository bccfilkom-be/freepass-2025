const { Session } = require("../models");

exports.getAllSessions = async (_req, res) => {
  try {
    const sessions = await Session.findAll({
      attributes: { exclude: ["createdAt", "updatedAt"] },
    });

    res.status(200).json(sessions);
  } catch (err) {
    res.status(500).json({ message: "Failed to retrieve sessions" });
  }
};
