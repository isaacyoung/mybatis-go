package base;

/**
 * 项目名称：auto-mybatis
 * 类描述：
 * 创建人：yzh
 * 创建时间：2017年06月01日
 * 备注：
 */
public class Tree {
    private String id;
    private String parent;
    private String text;
    private State state;

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getParent() {
        return parent;
    }

    public void setParent(String parent) {
        this.parent = parent;
    }

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }

    public void setOpened(boolean opened) {
        if (state == null) {
            state = new State();
        }
        this.state.setOpened(opened);
    }

    public void setDisabled(boolean disabled) {
        if (state == null) {
            state = new State();
        }
        this.state.setDisabled(disabled);
    }

    public void setSelected(boolean selected) {
        if (state == null) {
            state = new State();
        }
        this.state.setSelected(selected);
    }

    public State getState() {
        return state;
    }

    public void setState(State state) {
        this.state = state;
    }

    class State {
        private boolean opened = false;
        private boolean disabled = false;
        private boolean selected = false;

        public boolean isOpened() {
            return opened;
        }

        public void setOpened(boolean opened) {
            this.opened = opened;
        }

        public boolean isDisabled() {
            return disabled;
        }

        public void setDisabled(boolean disabled) {
            this.disabled = disabled;
        }

        public boolean isSelected() {
            return selected;
        }

        public void setSelected(boolean selected) {
            this.selected = selected;
        }
    }

}
