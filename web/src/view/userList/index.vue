<template>
  <div class="page">
    <div class="gva-card-box">
      <h1>111</h1>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@mixin flex-center {
  display: flex;
  align-items: center;
}
.page {
  background: #f0f2f5;
  padding: 0;
  .gva-card-box {
    padding: 12px 16px;
    & + .gva-card-box {
      padding-top: 0px;
    }
  }
  .gva-card {
    box-sizing: border-box;
    background-color: #fff;
    border-radius: 2px;
    height: auto;
    padding: 26px 30px;
    overflow: hidden;
    box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
  }
  .gva-top-card {
    height: 260px;
    @include flex-center;
    justify-content: space-between;
    color: #777;
    &-left {
      height: 100%;
      display: flex;
      flex-direction: column;
      &-title {
        font-size: 22px;
        color: #343844;
      }
      &-dot {
        font-size: 16px;
        color: #6b7687;
        margin-top: 24px;
      }
      &-rows {
        // margin-top: 15px;
        margin-top: 18px;
        color: #6b7687;
        width: 600px;
        align-items: center;
      }
      &-item {
        + .gva-top-card-left-item {
          margin-top: 24px;
        }
        margin-top: 14px;
      }
    }
    &-right {
      height: 600px;
      width: 600px;
      margin-top: 28px;
    }
  }
  ::v-deep(.el-card__header) {
    padding: 0;
    border-bottom: none;
  }
  .card-header {
    padding-bottom: 20px;
    border-bottom: 1px solid #e8e8e8;
  }
  .quick-entrance-title {
    height: 30px;
    font-size: 22px;
    color: #333;
    width: 100%;
    border-bottom: 1px solid #eee;
  }
  .quick-entrance-items {
    @include flex-center;
    justify-content: center;
    text-align: center;
    color: #333;
    .quick-entrance-item {
      padding: 16px 28px;
      margin-top: -16px;
      margin-bottom: -16px;
      border-radius: 4px;
      transition: all 0.2s;
      &:hover {
        box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
      }
      cursor: pointer;
      height: auto;
      text-align: center;
      // align-items: center;
      &-icon {
        width: 50px;
        height: 50px !important;
        border-radius: 8px;
        @include flex-center;
        justify-content: center;
        margin: 0 auto;
        i {
          font-size: 24px;
        }
      }
      p {
        margin-top: 10px;
      }
    }
  }
  .echart-box {
    padding: 14px;
  }
}
</style>
