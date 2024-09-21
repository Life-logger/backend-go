export const options = ({ url, method, headers, body }) => {
  return {
    url,
    method,
    headers,
    body,
    failOnStatusCode: false,
  };
};

export const todayAsString = () => {
  const now = new Date();
  const year = now.getFullYear();
  const month = now.getMonth() + 1;
  const date = now.getDate();
  return `${year}${appendZeroIfUnderTen(month)}${appendZeroIfUnderTen(date)}`;
};

function appendZeroIfUnderTen(n) {
  let str = n.toString();
  if (n < 10) {
    str = "0" + n;
  }
  return str;
}