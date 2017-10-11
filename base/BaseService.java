package base;

import com.github.pagehelper.Page;

import java.util.List;

/**
 * 项目名称：auto-mybatis
 * 类描述：
 * 创建人：yzh
 * 创建时间：2017年06月01日
 * 备注：
 */
public interface BaseService<T> {

    /**
     * 批量查询
     * @param bundle
     * @return
     */
    List<T> select(T bundle);

    /**
     * 单个查询
     * @param bundle
     * @return
     */
    T selectOne(T bundle);

    /**
     * 分页查询
     * @param bundle
     * @param page
     * @return
     */
    Page<T> selectByPage(T bundle, Page page);

    /**
     * 批量添加
     * @param bundle
     * @return
     */
    int insert(List<T> bundle);

    /**
     * 单个添加
     * @param bundle
     * @return
     */
    int insert(T bundle);

    /**
     * 批量更新
     * @param bundle
     * @return
     */
    int update(List<T> bundle);

    /**
     * 单个更新
     * @param bundle
     * @return
     */
    int update(T bundle);

    /**
     * 删除
     * @param bundle
     * @return
     */
    int delete(T bundle);

    /**
     * 根据id删除
     */
    int deleteById(Integer id);

    /**
     * 根据id查询
     */
    T selectById(Integer id);
}
