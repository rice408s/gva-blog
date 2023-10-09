<template>
    <div class="article-detail">
        <h2>{{ article.title }}</h2>
        <p>{{ article.content }}</p>
    </div>
</template>

<script>
export default {
    name: 'ArticleDetail',
    props: {
        article: {
            type: Object,
            required: true
        }
    }
};
</script>

<style scoped>
.article-detail {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
    background-color: #fff;
    border-radius: 4px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

h2 {
    font-size: 24px;
    font-weight: bold;
    margin-bottom: 20px;
}

p {
    font-size: 16px;
    line-height: 1.5;
}
</style>