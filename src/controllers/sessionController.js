const { Feedback, Session, User } = require("../models");

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

exports.leaveFeedback = async (req, res) => {
  const sessionId = req.params.id;
  const {
    body: { userId, feedback },
  } = req;

  try {
    const session = await Session.findByPk(sessionId);
    if (!session) {
      return res.status(404).json({ message: "Session not found" });
    }

    const user = await User.findByPk(userId);
    if (!user) {
      return res.status(404).json({ message: "User not found" });
    }

    const newFeedback = await Feedback.create({
      session_id: sessionId,
      user_id: userId,
      feedback: feedback,
    });

    res
      .status(201)
      .json({ message: "Feedback successfully subbmitted", newFeedback });
  } catch (err) {
    res.status(500).json({ message: "Failed to leave feedback" });
  }
};
