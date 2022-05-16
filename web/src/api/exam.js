import service from '@/utils/request.js';

export const getExamsList = (page = 1, size = 10) => service({
  url: `/exams?page=${page}&size=${size}`,
});

export const getExamsListByAdmin = (page = 1, size = 10) => service({
  url: `/admin/exams?page=${page}&size=${size}`,
});

export const getExam = (id) => service({
  url: `/exams/exam/${id}`,
});

export const getExamQuestions = () => service({
  url: '/exams/random',
});

export const getResult = (userExamId) => service({
  url: `/exams/answer/${userExamId}`,
});

export const submitAnswer = (postData) => service({
  url: '/exams/submit',
  method: 'post',
  data: postData,
});

export const delExam = (id) => service({
  url: `/admin/exams/exam/${id}`,
  method: 'delete',
});

export const editExam = (id, data) => service({
  url: `/admin/exams/exam/${id}`,
  method: 'put',
  data,
});

export const addExam = (data) => service({
  url: '/admin/exams/exam',
  method: 'post',
  data,
});
