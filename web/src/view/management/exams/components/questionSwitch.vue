<template>
  <el-switch v-model="value" @change="change"></el-switch>
</template>

<script>
export default {
  name: 'questionSwitch',
  data() {
    return {
      value: false,
    };
  },
  props: ['id', 'isChoosed'],
  watch: {
    isChoosed(newV) {
      this.value = newV;
    },
  },
  created() {
    this.value = this.isChoosed;
  },
  methods: {
    change(val) {
      const data = {
        id: this.id,
        isChoosed: val,
      };
      this.$emit('change', data);
    },
  },
};
</script>
