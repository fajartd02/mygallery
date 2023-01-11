// For database formatted datetime string
export function getFormattedDate(dbDateTimeFormatString) {
  let formattedDate = "";
  const date = new Date(dbDateTimeFormatString);

  const d = date.getDate(),
    m = date.getMonth() + 1,
    y = date.getFullYear();

  d < 10 ? (formattedDate += "0" + d + "/") : (formattedDate += d + "/");
  m < 10 ? (formattedDate += "0" + m + "/") : (formattedDate += m + "/");
  formattedDate += y;

  return formattedDate;
}

export function getFormattedTime(dbDateTimeFormatString) {
  const date = new Date(dbDateTimeFormatString);
  return date.toLocaleTimeString("en-US");
}
