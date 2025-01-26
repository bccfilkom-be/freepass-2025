const { Feedback, Session, SessionRegistration, User } = require("../models");

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

exports.registerForSession = async (req, res) => {
  const sessionId = req.params.id;
  const {
    body: { userId },
  } = req;

  try {
    const session = await Session.findByPk(sessionId);
    if (!session) {
      return res.status(404).json({ message: "Session not found" });
    }

    if (session.available_seats <= 0) {
      return res.status(400).json({ message: "No seats available" });
    }

    const existingRegistration = await SessionRegistration.findOne({
      where: { user_id: userId, session_id: sessionId },
    });
    if (existingRegistration) {
      return res
        .status(400)
        .json({ message: "Already registered for this session" });
    }

    await SessionRegistration.create({
      user_id: userId,
      session_id: sessionId,
    });
    session.available_seats -= 1;
    await session.save();

    res.status(201).json({ message: "Registration successful" });
  } catch (err) {
    console.error(err);
    res.status(500).json({ message: "Failed to register for session" });
  }
};
