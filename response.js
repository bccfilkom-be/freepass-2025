const response = (statusCode, data, message, res) => {
    res.status(statusCode).json({
        status_code: statusCode,
        datas: data,
        message: message        
    })
}

module.exports = response