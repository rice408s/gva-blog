<template>
    <div class="demo-page-header" style="background-color: #f5f5f5; padding: 24px">
        <a-page-header :ghost="ghost" :title="article.title" @back="() => $router.go(-1)">
            <!-- 副标题:sub-title="article.title" -->
            <template #extra>
                <a-button key="3">{{ article.title }}</a-button>
                <a-button key="2">Operation</a-button>
                <a-button key="1" type="primary">Primary</a-button>
            </template>
            <a-descriptions size="small" :column="3">
                <a-descriptions-item label="作者">{{ nickName ? nickName:'无名' }}</a-descriptions-item>
                <!-- <a-descriptions-item label="Association">
                    <a>421421</a>
                </a-descriptions-item> -->
                <a-descriptions-item label="发布时间">{{ formatDate(article.date) }}</a-descriptions-item>
                <a-descriptions-item label="字数统计" v-if="article.content">{{ article.content.length }}</a-descriptions-item>
                <!-- 阅读量 -->
                <!-- <a-descriptions-item label="阅读量">{{ article.readCount }}</a-descriptions-item> -->
                <!-- 评论数 -->
                <!-- <a-descriptions-item label="评论数">{{ article.commentCount }}</a-descriptions-item> -->
            </a-descriptions>
        </a-page-header>
    </div>
    <!-- <a-checkbox v-model:checked="ghost" style="margin-top: 0.5rem">
    toggle ghost</a-checkbox> -->
</template>
<script >
import { ref ,watchEffect} from 'vue';
import axios from 'axios';
import { formatDate } from '@/utils/time';
// import { onMounted } from 'vue';
// const article=ref()

export default {
    name: 'PageHeader',
    props: {
        article: {
            type: Object,
            default: () => { }
        },
        userId:{
            type:Number,
            default:()=>{}
        }
    },

    methods: {
        formatDate
    }, 

    setup(props) {
        // const article=ref({title:'title'})
        //接受父组件传递的article
        const nickName = ref('')


        //监听userId的变化
        watchEffect(async () => {
            try {
                //直接使用props.userId,如果使用article.value.userId刷新会空指针
                const res = await axios.get("/user/getUsernameById?ID=" +props.userId);
                nickName.value = res.data.data.nickName
                
                console.log('1' + nickName.value)
            } catch (e) {
                console.log(e)
            }
        });


        //监听article的变化
        // watchEffect(() => {
        //     console.log(article.value)
        // })

        return {
            nickName
        }

    }
};



</script>



<style scoped>
.demo-page-header :deep(tr:last-child td) {
    padding-bottom: 0;
    /*  */
}

.demo-page-header {
    padding-top: 0;
}
</style>