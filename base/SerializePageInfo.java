package base;

import com.github.pagehelper.Page;

/**
 * 项目名称：auto-mybatis
 * 类描述：序列化 分页信息
 * 创建人：yzh
 * 创建时间：2017年06月01日
 * 备注：
 */
public class SerializePageInfo {

    private int pageNum;
    private int pageSize;
    private long total;

    public SerializePageInfo(Page page) {
        pageNum=page.getPageNum();
        pageSize=page.getPageSize();
        total=page.getTotal();
    }

    public int getPageNum() {
        return pageNum;
    }

    public void setPageNum(int pageNum) {
        this.pageNum = pageNum;
    }

    public int getPageSize() {
        return pageSize;
    }

    public void setPageSize(int pageSize) {
        this.pageSize = pageSize;
    }

    public long getTotal() {
        return total;
    }

    public void setTotal(long total) {
        this.total = total;
    }
}
