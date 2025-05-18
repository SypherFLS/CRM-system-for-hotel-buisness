<template>
    <div id="app" class="body">
        <div class="top-part">
            <div class="toper-container">
                <img class="background-pic" src="@/assets/hotel_pic.png" alt="hotel" />
            </div>
            
            <Header></Header>
        </div>
    </div>
</template>

<script>
import Header from './components/Header.vue'; 
    export default {
        components: {
            Header
        }
    }
</script>

<style scoped>
    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }
    
    .top-part {
        margin: 0;
        width: 100%;
        height: 100vh; 
        overflow: hidden;
        position: relative;
    }
    
    

    .background-pic {
        position: absolute;
        top: 0;
        left: 0;
        width: 110%;
        height: auto;
        max-height: 500px;
        object-fit: cover; 
        z-index: -1; 
    }

    .header {
        z-index: 1; 
    }
</style>