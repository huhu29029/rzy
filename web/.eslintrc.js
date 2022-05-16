module.exports = {
  root: true,
  extends: [
    'airbnb-base',
    'plugin:vue/essential',
  ],
  // add your custom rules here
  rules: {
    'no-extend-native': ['warn'],
    'no-console': ['error', { allow: ['warn', 'error'] }],
    'import/no-unresolved': ['off'],
    'import/extensions': ['off'],
    'import/no-extraneous-dependencies': ['off'],
    'no-plusplus': ['off'],
    'no-param-reassign': ['off'],
    'prefer-destructuring': ['off'],
    'no-shadow': ['off'],
    'max-len': ['error', 150],
    'import/no-named-as-default': ['off'],
    'import/no-named-as-default-member': ['off'],
    'import/prefer-default-export': ['off'],
  },
};
