<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容:" prop="content">
          <el-input v-model="formData.content" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发布日期:" prop="date">
          <el-date-picker v-model="formData.date" type="date" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="点赞数:" prop="likes">
          <el-input v-model.number="formData.likes" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标签:" prop="tags">
          <el-input v-model="formData.tags" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类:" prop="category">
          <el-input v-model="formData.category" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item, key) in ArticleOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { createArticle, updateArticle, findArticle } from "@/api/article";

defineOptions({
  name: "ArticleForm",
});

// 自动获取字典
import { getDictFunc } from "@/utils/format";
import { useRoute, useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { ref, reactive } from "vue";
import { useUserStore } from "src/pinia/modules/user.js";
const route = useRoute();
const router = useRouter();

console.log(useUserStore);
const type = ref("");
const ArticleOptions = ref([]);
const formData = ref({
  title: "",
  userId: 0,
  content: "",
  date: new Date(),
  likes: 0,
  tags: "",
  category: "",
  status: undefined,
});
// 验证规则
const rule = reactive({
  title: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  userId: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  content: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  date: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  likes: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  status: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
});

const elFormRef = ref();

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findArticle({ ID: route.query.id });
    if (res.code === 0) {
      formData.value = res.data.rearticle;
      type.value = "update";
    }
  } else {
    type.value = "create";
  }
  ArticleOptions.value = await getDictFunc("Article");
};

init();
// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createArticle(formData.value);
        break;
      case "update":
        res = await updateArticle(formData.value);
        break;
      default:
        res = await createArticle(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
    }
  });
};

// 返回按钮
const back = () => {
  router.go(-1);
};
</script>

<style></style>
