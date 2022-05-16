import service from '@/utils/request.js';

export const getRandomQuestion = () => service({
  url: '/questions/random',
});

export const getChapterRandomQuestion = (ChapterId) => service({
  url: '/questions/random',
  method: 'post',
  data: ChapterId,
});

// 算法推荐题目
export const getQuestionWithKT = () => service({
  url: '/questions/kt',
});

// 错题
export const getQuestionWithWrong = () => service({
  url: '/questions/wrong',
});
export const getTags = () => service({
  url: '/tags',
});

export const getChapterId = () => service({
  url: '/chapters',
});

export const submitAnswer = (postData) => service({
  url: '/questions/submit',
  method: 'post',
  data: postData,
});

export const getQuestions = (page = 1, size = 10) => service({
  // url: `/questions?page=${page}&size=${size}`
  url: `/admin/questions?page=${page}&size=${size}`,
});

export const getQuestionsHistory = (page = 1, size = 10) => service({
  url: `/user/questions?page=${page}&size=${size}`,
});

export const getExamHistory = (page = 1, size = 10) => service({
  url: `/user/exams?page=${page}&size=${size}`,
});

export const getUserInfoExamHistory = (id = 0, page = 1, size = 10) => service({
  url: `/admin/users/${id}/exams?page=${page}&size=${size}`,
});

export const getUserInfoQuestionsHistory = (id = 0, page = 1, size = 10) => service({
  url: `/admin/users/${id}/questions?page=${page}&size=${size}`,
});

export const getQuestionResult = (id) => service({
  url: `/questions/answer/${id}`,
});

export const getQuestion = (id) => service({
  url: `/questions/question/${id}`,
});

export const addQuestion = (data) => service({
  // url: '/questions',
  url: '/admin/questions',
  method: 'post',
  data,
});

export const editQuestion = (id, data) => service({
  // url: `/questions/${id}`,
  url: `/admin/questions/${id}`,
  method: 'put',
  data,
});

export const delQuestion = (id) => service({
  // url: `/questions/${id}`,
  url: `/admin/questions/${id}`,
  method: 'delete',
});

export const getWrongQuestion = (page = 1, size = 10) => service({
  url: `/user/WrongQuestion?page=${page}&size=${size}`,
});

export const delWrongQuestion = (id) => service({
  url: `/user/DeleteQuestionFromWrongQuestionBook?id=${id}`,
  method: 'delete',
});
