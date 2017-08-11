package base;

import java.util.Set;

/**
 * 项目名称：auto-mybatis
 * 类描述：数据权限
 * 创建人：yzh
 * 创建时间：2017/6/13
 * 备注：
 */
public class DataRightResult {
    /**
     * 所有权限
     */
    private boolean allRights;
    /**
     * 权限 部门
     */
    private Set<Integer> deptIds;
    /**
     * 权限 用户
     */
    private Set<Integer> users;

    private Integer noneRight;

    /**
     * 没有权限
     * @return
     */
    public boolean isEmpty() {
        return !allRights && (deptIds == null || deptIds.isEmpty()) && (users == null || users.isEmpty());
    }

    public boolean isAllRights() {
        return allRights;
    }

    public boolean hasDeptRights() {
        return deptIds != null && !deptIds.isEmpty();
    }

    public boolean hasUserRights() {
        return users != null && !users.isEmpty();
    }

    public void setAllRights(boolean allRights) {
        this.allRights = allRights;
        return ;
    }

    public Set<Integer> getDeptIds() {
        return deptIds;
    }

    public void setDeptIds(Set<Integer> deptIds) {
        this.deptIds = deptIds;
    }

    public Set<Integer> getUsers() {
        return users;
    }

    public void setUsers(Set<Integer> users) {
        this.users = users;
    }

    public Integer getNoneRight() {
        return noneRight;
    }

    public void setNoneRight(Integer noneRight) {
        this.noneRight = noneRight;
    }

    public void setDataRight(DataRightResult dataRightResult) {
        setNoneRight(dataRightResult.getNoneRight());
        setDeptIds(dataRightResult.getDeptIds());
        setUsers(dataRightResult.getUsers());
    }
}
