import service from '@/utils/request.js';

function newXAPIData(data = {
  // actor: {},
  verb: '',
  object: {
    info: '',
  },
  // result: {},
  // context: {},
  // authority: {},
  // timestamp: new Date().toISOString()
}) {
  return service({
    url: '/xapi/new',
    method: 'post',
    data,
  });
}

export {
  newXAPIData,
};
