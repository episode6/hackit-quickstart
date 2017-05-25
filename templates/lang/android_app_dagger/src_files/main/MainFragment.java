package {{ .Group }}.main;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import {{ .Group }}.R;
import {{ .Group }}.base.BaseFragment;

import javax.annotation.Nullable;

import butterknife.BindView;
import butterknife.ButterKnife;

public class MainFragment extends BaseFragment {

  @BindView(R.id.tv_hello_world) TextView mTextView;

  @Nullable
  @Override
  public View onCreateView(
      LayoutInflater inflater,
      @Nullable ViewGroup container,
      Bundle savedInstanceState) {
    return inflater.inflate(R.layout.fragment_main, container, false);
  }


  @Override
  public void onViewCreated(final View view, @Nullable Bundle savedInstanceState) {
    ButterKnife.bind(this, view);
    mTextView.setText("Hello Earth!");
  }
}
