function objToMap(obj) {
  const map = new Map();
  // eslint-disable-next-line
  for (const k of Object.keys(obj)) {
    map.set(Number(k), obj[k]);
  }
  return map;
}
function mapToObj(map) {
  const obj = Object.create(null);
  // eslint-disable-next-line
  for (const [k, v] of map) {
    obj[k] = v;
  }
  return obj;
}

export {
  objToMap,
  mapToObj,
};
