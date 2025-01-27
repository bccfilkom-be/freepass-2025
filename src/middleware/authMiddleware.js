const jwt = require('jsonwebtoken');
require('dotenv').config();

const authMiddleware = (roles = []) => {
    if (typeof roles === 'string') {
        roles = [roles];
    }

    return (req, res, next) => {
        try {
            const token = req.headers.authorization;

            if (!token) {
                return res.status(401).json({
                    error: true,
                    code: 401,
                    message: 'Unauthorized',
                    description: 'Token tidak ditemukan!',
                });
            }

            const tokenWithoutBearer = token.split(' ')[1];
            const key = process.env.API_KEY;

            jwt.verify(tokenWithoutBearer, key, (err, decoded) => {
                if (err) {
                    if (err.name === 'TokenExpiredError') {
                        return res.status(401).json({
                            error: true,
                            code: 401,
                            message: 'Unauthorized',
                            description: 'Token telah kedaluwarsa!',
                        });
                    } else {
                        return res.status(401).json({
                            error: true,
                            code: 401,
                            message: 'Unauthorized',
                            description: 'Token tidak valid!',
                        });
                    }
                } else {
                    req.user = decoded;
                    if (roles.length && !roles.includes(decoded.user_role)) {
                        return res.status(403).json({
                            error: true,
                            code: 403,
                            message: 'Forbidden',
                            description: 'Anda tidak memiliki akses ke resource ini.',
                        });
                    }

                    next();
                }
            });
        } catch (error) {
            return res.status(500).json({
                error: true,
                code: 500,
                message: 'Internal Server Error',
                description: 'Terjadi kesalahan pada server.',
                errorMessage: error.message,
            });
        }
    };
};

module.exports = authMiddleware;
