<template>    
  <div class="home">

      <!-- <h1 class="page-title">文章列表</h1> -->
      <div class="article-list">
        <div class="article-item" v-for="article in articles" :key="article.ID" @click="toArticleDetail(article.ID)">
          <!-- <router-link :to="'/article/' + article.id"> -->
          <div class="article-header">  
            <h3 class="article-title"> {{ article.title }}</h3>
            <p class="article-meta">
              <span class="article-category">{{ article.category }}</span>
              <span class="article-date">{{ formatDate(article.date) }}</span>
            </p>
          </div>
          <!-- </router-link> -->
          <p class="article-content">{{ article.content }}</p>
        </div>



  </div>


</div></template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

import router from '../router/index'

const articles = ref([]);




onMounted(async () => {
  try {
    const res = await axios.get('/article/getArticleList');
    articles.value = res.data.data.list;
  } catch (error) {
    console.error(error);
  }
});


  // const hello= ()=>{
  //   router.push('/login')
  // }


const toArticleDetail = (id) => {
  console.log(id)
  router.push({name: 'ArticleDetail', params: {id}});
};


const formatDate = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  });
};

</script>

<style scoped>
.home {
  /* padding-top: 80px; */
  width: 100%;
  padding-top: 40px;
  background-image: "../assets/backImage.jpeg";
  background-size: cover;
  justify-content: center;
  /* overflow-y: scroll; */
  /* padding-right: calc(100% - 100vw); /* 假设滚动条宽度为默认值 */
  /* ::-webkit-scrollbar {
  width: 16px; 假设滚动条宽度为16px */ 
/* } */
}



.page-title {
  font-size: 36px;
  font-weight: bold;
  /* margin-bottom: 10px; */
  margin-top: 30px;
}

.article-list {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 200%;
  max-width: 600px;
  padding-top: 350px;


}

.article-item {
  width: 200%;
  max-width: 800px;
  margin-bottom: 20px;
  padding: 20px;
  box-sizing: border-box;
  background-color: #fff;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  border-radius: 4px;
  transition: all 0.3s ease;
  margin-left: 0;
  margin-right: 0;
}

.article-item:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  transform: translateY(-2px);
}

.article-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.article-title {
  font-size: 24px;
  font-weight: bold;
  margin: 0;
}

.article-meta {
  font-size: 14px;
  color: #999;
  margin: 0;
}

.article-date {
  margin-right: 10px;
}

.article-category {
  padding: 4px 8px;
  border-radius: 4px;
  background-color: #eee;
}

.article-content {
  font-size: 18px;
  line-height: 1.5;
  overflow: hidden;
  text-overflow: ellipsis;
}


.el-pagination {
  display: flex;
  justify-content: center;
  margin-top: 20px;
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: #fff;
  box-shadow: 0 -2px 12px rgba(0, 0, 0, 0.1);
  padding: 10px;
}




</style>