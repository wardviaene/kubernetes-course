module.exports = {
  uppercase: function (event, context) {
    str = event['data'].toUpperCase()
    console.log(str);
    return str
  }
}
