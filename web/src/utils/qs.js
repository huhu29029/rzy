const qs = require('qs');

const parseHash = (hash) => {
  // hash: "#/login?redirect=%2Flayout%2Fexam%2Fguide"
  // query: "redirect=%2Flayout%2Fexam%2Fguide"
  // queryObj: {redirect: "/layout/exam/guide"}
  const query = hash.split('?')[1];
  const queryObj = qs.parse(query);
  return queryObj;
};

export {
  parseHash,
};
