<template>
    <div class="footer">
        <div class="footer-content">
            <div class="footer-text">
                © 2021 My Website. All rights reserved.
            </div>
        </div>
    </div>
</template>

<style scoped>
.footer {
    /* background-color: #d6e1f3; */
    height: 60px;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 185%;
    margin-left: 0;
}

.footer-content {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    width: 200%;
    max-width: 1400px;
    /* padding: ; */
    /* margin: 0; */
    /* border: 0px solid #000; */
}

.footer-links {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
}

.footer-links a {
    color: #666;
    margin-right: 16px;
}

.footer-links a:last-child {
    margin-right: 0;
}

.footer-text {
    color: #666;
}
</style>