import service from '@/utils/request.js';

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => service({
  url: '/base/login',
  method: 'post',
  data,
});

// @Summary 获取验证码
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/captcha [post]
export const captcha = (data) => service({
  url: '/base/captcha',
  method: 'post',
  data,
});

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/resige [post]
export const register = (data) => service({
  url: '/base/register',
  method: 'post',
  data,
});

// @Summary 修改密码
// @Produce  application/json
// @Param data body {username:"string",password:"string",newPassword:"string"}
// @Router /user/changePassword [post]
export const changePassword = (data) => service({
  url: '/user/changePassword',
  method: 'post',
  data,
});

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserInfo [put]
export const setUserInfo = (data) => service({
  url: '/user/setUserInfo',
  method: 'put',
  data,
});

export const getUserInfo = () => service({
  url: '/user',
});

export const getUserList = (page = 1, size = 10) => service({
  url: `/admin/users?page=${page}&size=${size}`,
});
