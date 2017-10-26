package tech.evove.goandroid.core;


import io.reactivex.disposables.Disposable;

final class GoDisposable implements core.Disposable, Disposable {
    private final core.Disposable actual;

    private GoDisposable(core.Disposable actual) {
        this.actual = actual;
    }

    static Disposable wrap(core.Disposable disposable) {
        if (disposable instanceof Disposable) {
            return (Disposable) disposable;
        }
        return new GoDisposable(disposable);
    }

    @Override
    public void dispose() {
        actual.dispose();
    }

    @Override
    public boolean isDisposed() {
        return actual.isDisposed();
    }
}
