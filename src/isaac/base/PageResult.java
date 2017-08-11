package base;

import java.util.List;

/**
 * 项目名称：auto-mybatis
 * 类描述：
 * 创建人：yzh
 * 创建时间：2017年06月01日
 * 备注：
 */
public class PageResult {
    private long total;
    private List rows;

    public long getTotal() {
        return total;
    }

    public void setTotal(long total) {
        this.total = total;
    }

    public List getRows() {
        return rows;
    }

    public void setRows(List rows) {
        this.rows = rows;
    }
}
