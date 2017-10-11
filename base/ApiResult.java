package base;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.serializer.SimplePropertyPreFilter;
import com.github.pagehelper.Page;

/**
 * 项目名称：auto-mybatis
 * 类描述：JSON返回
 * 创建人：yzh
 * 创建时间：2017年06月01日
 * 备注：
 */
public class ApiResult<T> {
    public static final String SUCCESS = "请求成功";
    public static final int SUCCESS_STATUS_CODE = 0;

    public static final String FAILURE = "请求失败";
    public static final int FAILURE_STATUS_CODE = -1;

    /**
     * success:0  failure:-1
     */
    private int code;
    /**
     * 消息
     */
    private String msg;
    /**
     * 数据
     */
    private T data;

    /**
     * 分页信息
     */
    private SerializePageInfo page;

    public int getCode() {
        return code;
    }

    public void setCode(int code) {
        this.code = code;
    }

    public String getMsg() {
        return msg;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }

    public T getData() {
        return data;
    }

    public void setData(T data) {
        this.data = data;
    }


    public SerializePageInfo getPage() {
        return page;
    }

    public void setPage(SerializePageInfo page) {
        this.page = page;
    }

    public String toJSONString(Class<T> clazz, String[] filterArr) {
        SimplePropertyPreFilter filter = new SimplePropertyPreFilter(clazz, filterArr);
        return JSON.toJSONString(this, filter);
    }

    public String toJSONString(SimplePropertyPreFilter filter) {
        return JSON.toJSONString(this, filter);
    }

    public String toJSONString(String[] filter) {
        return JSON.toJSONString(this, new SimplePropertyPreFilter(filter));
    }

    public String toJSONString() {
        return JSON.toJSONString(this);
    }

    public ApiResult<T> returnFailureResult() {
        this.setData(null);
        this.setCode(FAILURE_STATUS_CODE);
        this.setMsg(FAILURE);
        return this;
    }

    public ApiResult<T> returnFailureResult(T o) {
        this.setData(o);
        this.setCode(FAILURE_STATUS_CODE);
        this.setMsg(FAILURE);
        return this;
    }

    public ApiResult<T> returnFailureResult(String msg) {
        this.setData(null);
        this.setCode(FAILURE_STATUS_CODE);
        this.setMsg(msg);
        return this;
    }


    public ApiResult<T> returnSuccessResult(T o) {
        this.setData(o);
        this.setCode(SUCCESS_STATUS_CODE);
        this.setMsg(SUCCESS);
        return this;
    }

    public ApiResult<T> returnSuccessResult(T o,Page page) {
        this.setData(o);
        this.setCode(SUCCESS_STATUS_CODE);
        this.setMsg(SUCCESS);
        this.setPage(new SerializePageInfo(page));
        return this;
    }

    public ApiResult<T> returnSuccessResult(String o) {
        this.setData(null);
        this.setCode(SUCCESS_STATUS_CODE);
        this.setMsg(o);
        return this;
    }

    public static<T> ApiResult<T> newInstance(){
        return new ApiResult<T>();
    }
}
