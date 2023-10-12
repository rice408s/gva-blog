<template>
    <div class="article-detail">
        <el-card>
            <h1 slot="header">{{ article.title }}</h1>
            <div class="article-meta">
                <span>作者：{{ article.author }}</span>
                <span>发布时间：{{ article.publishTime }}</span>
                <span>阅读量：{{ article.readCount }}</span>
            </div>
            <div class="article-content" v-html="article.content"></div>
            <div class="article-comments">
                <h2>评论区</h2>
                <el-comment v-for="comment in comments" :key="comment.id" :author="comment.author" :datetime="comment.time">
                    <div slot="content">{{ comment.content }}</div>
                    <el-comment v-for="reply in comment.replies" :key="reply.id" :author="reply.author" :datetime="reply.time">
                        <div slot="content">{{ reply.content }}</div>
                    </el-comment>
                    <el-form-item>
                        <el-input v-model="comment.replyContent" placeholder="请输入回复内容"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="submitReply(comment)">回复</el-button>
                    </el-form-item>
                </el-comment>
                <el-form>
                    <el-form-item>
                        <el-input v-model="commentContent" placeholder="请输入评论内容"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="submitComment">发表评论</el-button>
                    </el-form-item>
                </el-form>
            </div>
            <div class="article-actions">
                <el-button type="primary" icon="el-icon-share" @click="share">分享</el-button>
                <el-button type="primary" icon="el-icon-thumb-up" @click="like">点赞</el-button>
                <el-button type="primary" icon="el-icon-star-on" @click="collect">收藏</el-button>
            </div>
            <div class="article-related">
                <h2>相关博客推荐</h2>
                <el-row>
                    <el-col v-for="related in relatedArticles" :key="related.id" :span="8">
                        <el-card :body-style="{ padding: '0px' }">
                            <img :src="related.image" class="related-image">
                            <div style="padding: 14px;">
                                <h4>{{ related.title }}</h4>
                                <p>{{ related.summary }}</p>
                            </div>
                        </el-card>
                    </el-col>
                </el-row>
            </div>
        </el-card>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';

export default {
    name: 'ArticleDetail',
    props: {
        id: {
            type: String,
            required: true,
        },
    },
    setup(props) {
        const article = ref({});
        const comments = ref([]);
        const commentContent = ref('');
        const relatedArticles = ref([]);

        onMounted(async () => {
            try {
                // 获取博客详情
                const res = await axios.get(`/article/findArticle?ID=${props.id}`);
                article.value = res.data.data.rearticle;

                // 获取评论列表
                const commentRes = await axios.get(`/comment/list?articleId=${props.id}`);
                comments.value = commentRes.data.data;

                // 获取相关博客推荐
                const relatedRes = await axios.get(`/article/related?articleId=${props.id}`);
                relatedArticles.value = relatedRes.data.data;
            } catch (error) {
                console.error(error);
            }
        });

        // 发表评论
        const submitComment = async () => {
            try {
                const res = await axios.post('/comment/create', {
                    articleId: props.id,
                    content: commentContent.value,
                });
                comments.value.push(res.data.data);
                commentContent.value = '';
            } catch (error) {
                console.error(error);
            }
        };

        // 回复评论
        const submitReply = async (comment) => {
            try {
                const res = await axios.post('/comment/reply', {
                    commentId: comment.id,
                    content: comment.replyContent,
                });
                comment.replies.push(res.data.data);
                comment.replyContent = '';
            } catch (error) {
                console.error(error);
            }
        };

        // 分享博客
        const share = () => {
            // TODO: 实现分享功能
        };

        // 点赞博客
        const like = () => {
            // TODO: 实现点赞功能
        };

        // 收藏博客
        const collect = () => {
            // TODO: 实现收藏功能
        };

        return {
            article,
            comments,
            commentContent,
            relatedArticles,
            submitComment,
            submitReply,
            share,
            like,
            collect,
        };
    },
};
</script>

<style scoped>
.article-detail {
    max-width: 100%;
    margin: 0 auto;
    padding: 20px;
}

.article-meta {
    margin-bottom: 20px;
}

.article-meta span {
    margin-right: 20px;
}

.article-content {
    margin-bottom: 20px;
}

.article-comments {
    margin-bottom: 20px;
}

.comment-list {
    list-style: none;
    padding: 0;
}

.comment-list li {
    margin-bottom: 20px;
}

.comment-header {
    margin-bottom: 10px;
}

.comment-header span {
    margin-right: 10px;
}

.comment-content {
    margin-bottom: 10px;
}

.reply-list {
    list-style: none;
    padding: 0;
    margin-top: 10px;
    margin-bottom: 10px;
}

.reply-list li {
    margin-bottom: 10px;
}

.reply-header {
    margin-bottom: 5px;
}

.reply-header span {
    margin-right: 10px;
}

.reply-content {
    margin-bottom: 5px;
}

.reply-form {
    margin-top: 10px;
    margin-bottom: 10px;
}

.reply-form textarea {
    width: 100%;
    height: 60px;
    margin-bottom: 10px;
}

.comment-form textarea {
    width: 100%;
    height: 80px;
    margin-bottom: 10px;
}

.article-actions {
    margin-bottom: 20px;
}

.article-actions button {
    margin-right: 10px;
}

.article-related {
    margin-bottom: 20px;
}

.article-related h2 {
    margin-bottom: 10px;
}

.related-list {
    list-style: none;
    padding: 0;
}

.related-list li {
    margin-bottom: 10px;
}

.related-list a {
    color: #333;
    text-decoration: none;
}

.related-image {
    width: 100%;
    height: 200px;
    object-fit: cover;
}
</style>