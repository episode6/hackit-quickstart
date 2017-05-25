package {{ .Group }}.main;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import {{ .Group }}.R;
import {{ .Group }}.base.BaseFragment;

import javax.annotation.Nullable;

public class MainFragment extends BaseFragment {

  @Nullable
  @Override
  public View onCreateView(
      LayoutInflater inflater,
      @Nullable ViewGroup container,
      Bundle savedInstanceState) {
    return inflater.inflate(R.layout.fragment_main, container, false);
  }
}
